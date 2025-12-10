const std = @import("std");
const ArrayList = std.ArrayList;

pub fn run(alloc: std.mem.Allocator) !void {
    std.debug.print("Day 8\n", .{});
    const part1 = try task1(alloc, comptime @embedFile("inputs/day8_input.txt"), 1000);
    std.debug.print("Part 1: {d}\n", .{part1});
    const part2 = try task2(alloc, comptime @embedFile("inputs/day8_input.txt"));
    std.debug.print("Part 2: {d}\n", .{part2});
}

const Point = struct { x: u32, y: u32, z: u32 };

fn parseInput(alloc: std.mem.Allocator, input: []const u8) !ArrayList(Point) {
    var list = ArrayList(Point){};
    var line_iter = std.mem.tokenizeScalar(u8, input, '\n');

    while (line_iter.next()) |line| {
        var num_iter = std.mem.tokenizeScalar(u8, line, ',');
        const p = Point{
            .x = try std.fmt.parseInt(u32, num_iter.next().?, 10),
            .y = try std.fmt.parseInt(u32, num_iter.next().?, 10),
            .z = try std.fmt.parseInt(u32, num_iter.next().?, 10),
        };
        try list.append(alloc, p);
    }

    return list;
}

fn calculateDistance(p1: Point, p2: Point) f64 {
    const dx = @as(f64, @floatFromInt(@max(p1.x, p2.x) - @min(p1.x, p2.x)));
    const dy = @as(f64, @floatFromInt(@max(p1.y, p2.y) - @min(p1.y, p2.y)));
    const dz = @as(f64, @floatFromInt(@max(p1.z, p2.z) - @min(p1.z, p2.z)));
    const dist = @sqrt(dx * dx + dy * dy + dz * dz);
    return dist;
}

const PointDistance = struct {
    p1: usize,
    p2: usize,
    dist: f64,
};

fn comparePointDistanceAscending(_: void, p1: PointDistance, p2: PointDistance) bool {
    return p1.dist > p2.dist;
}

fn comparePointDistanceDescending(_: void, p1: PointDistance, p2: PointDistance) bool {
    return p1.dist < p2.dist;
}

fn calculateAllDistances(alloc: std.mem.Allocator, points: ArrayList(Point), ascending: bool) !ArrayList(PointDistance) {
    var distances = ArrayList(PointDistance){};

    for (0..points.items.len) |i| {
        for (i + 1..points.items.len) |j| {
            const dist = calculateDistance(points.items[i], points.items[j]);
            try distances.append(alloc, PointDistance{
                .p1 = i,
                .p2 = j,
                .dist = dist,
            });
        }
    }

    if (ascending) {
        std.mem.sortUnstable(PointDistance, distances.items, {}, comparePointDistanceAscending);
    } else {
        std.mem.sortUnstable(PointDistance, distances.items, {}, comparePointDistanceDescending);
    }
    return distances;
}

fn buildAdjacencyList(alloc: std.mem.Allocator, distances: *ArrayList(PointDistance), connections: u32) !std.AutoHashMap(usize, ArrayList(usize)) {
    var adj_list = std.AutoHashMap(usize, ArrayList(usize)).init(alloc);

    for (0..connections) |_| {
        const closest = distances.pop().?;
        var list1 = try adj_list.getOrPutValue(closest.p1, ArrayList(usize){});
        try list1.value_ptr.append(alloc, closest.p2);
        var list2 = try adj_list.getOrPutValue(closest.p2, ArrayList(usize){});
        try list2.value_ptr.append(alloc, closest.p1);
    }

    return adj_list;
}

fn findConnectedComponent(alloc: std.mem.Allocator, start: usize, adj_list: std.AutoHashMap(usize, ArrayList(usize)), visited: *std.AutoHashMap(usize, void)) !u32 {
    var size: u32 = 0;
    var stack = ArrayList(usize){};
    defer stack.deinit(alloc);
    try stack.append(alloc, start);

    while (stack.items.len != 0) {
        const node = stack.pop().?;
        size += 1;
        const neighbors = adj_list.get(node);
        if (neighbors == null) continue;
        for (neighbors.?.items) |nb| {
            if (!visited.contains(nb)) {
                try visited.put(nb, {});
                try stack.append(alloc, nb);
            }
        }
    }

    return size;
}

fn findAllComponentSizes(alloc: std.mem.Allocator, point_count: usize, adj_list: std.AutoHashMap(usize, ArrayList(usize))) !ArrayList(u32) {
    var sizes = ArrayList(u32){};
    var visited = std.AutoHashMap(usize, void).init(alloc);
    defer visited.deinit();

    for (0..point_count - 1) |i| {
        if (visited.contains(i)) continue;
        try visited.put(i, {});

        const size = try findConnectedComponent(alloc, i, adj_list, &visited);
        try sizes.append(alloc, size);
    }

    std.mem.sort(u32, sizes.items, {}, std.sort.desc(u32));
    return sizes;
}

fn task1(alloc: std.mem.Allocator, input: []const u8, connections: u32) !u64 {
    var points = try parseInput(alloc, input);
    defer points.deinit(alloc);

    var distances = try calculateAllDistances(alloc, points, true);
    defer distances.deinit(alloc);

    var adj_list = try buildAdjacencyList(alloc, &distances, connections);
    defer {
        var val_iter = adj_list.valueIterator();
        while (val_iter.next()) |item| {
            item.deinit(alloc);
        }
        adj_list.deinit();
    }

    var sizes = try findAllComponentSizes(alloc, points.items.len, adj_list);
    defer sizes.deinit(alloc);

    var res: u64 = 1;
    res *= sizes.items[0];
    res *= sizes.items[1];
    res *= sizes.items[2];

    return res;
}

fn resolveCircuit(p: usize, circuit_mapping: []usize, circuits: std.AutoHashMap(usize, void)) usize {
    var circuit = circuit_mapping[p];
    while (!circuits.contains(circuit)) {
        circuit = circuit_mapping[circuit];
    }
    return circuit;
}

fn task2(alloc: std.mem.Allocator, input: []const u8) !u64 {
    var points = try parseInput(alloc, input);
    defer points.deinit(alloc);
    const num_points = points.items.len;

    var distances = try calculateAllDistances(alloc, points, false);
    defer distances.deinit(alloc);

    // Set of current circuits
    var circuits = std.AutoHashMap(usize, void).init(alloc);
    defer circuits.deinit();
    // Original Circuit: Current Circuit
    var circuit_mapping = try alloc.alloc(usize, num_points);
    defer alloc.free(circuit_mapping);
    for (0..points.items.len) |i| {
        try circuits.put(i, {});
        circuit_mapping[i] = i;
    }

    // Iterate over distances
    for (distances.items) |d| {
        const p1 = d.p1;
        const p2 = d.p2;
        const c1 = resolveCircuit(p1, circuit_mapping, circuits);
        const c2 = resolveCircuit(p2, circuit_mapping, circuits);
        if (c1 == c2) continue;
        if (c1 < c2) {
            // Change c2 to c1
            circuit_mapping[c2] = c1;
            _ = circuits.remove(c2);
        } else {
            // change c1 to c2
            circuit_mapping[c1] = c2;
            _ = circuits.remove(c1);
        }
        if (circuits.count() == 1) {
            var ans: u64 = 1;
            ans *= points.items[p1].x;
            ans *= points.items[p2].x;
            return ans;
        }
    }

    return 0;
}

test "task1" {
    const input = comptime @embedFile("inputs/day8_sample.txt");
    const result = try task1(std.testing.allocator, input, 10);
    try std.testing.expectEqual(result, 40);
}

test "task2" {
    const input = comptime @embedFile("inputs/day8_sample.txt");
    const result = try task2(std.testing.allocator, input);
    try std.testing.expectEqual(result, 25272);
}
