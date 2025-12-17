const std = @import("std");
const ArrayList = std.ArrayList;
const Queue = @import("common/queue.zig").Queue;

const Device = [3]u8;

pub fn run(alloc: std.mem.Allocator) !void {
    std.debug.print("Day 11\n", .{});
    const part1 = try task1(alloc, comptime @embedFile("inputs/day11_input.txt"));
    std.debug.print("Part 1: {d}\n", .{part1});
    const part2 = try task2(alloc, comptime @embedFile("inputs/day11_input.txt"));
    std.debug.print("Part 2: {d}\n", .{part2});
}

fn parseGraph(alloc: std.mem.Allocator, input: []const u8) !std.AutoHashMap([3]u8, ArrayList([3]u8)) {
    var line_iter = std.mem.tokenizeScalar(u8, input, '\n');
    var graph = std.AutoHashMap([3]u8, ArrayList([3]u8)).init(alloc);

    while (line_iter.next()) |line| {
        var iter = std.mem.tokenizeScalar(u8, line, ' ');
        const key_slice = iter.next().?[0..3];
        var key: Device = undefined;
        @memcpy(&key, key_slice);

        var vals = ArrayList(Device){};
        while (iter.next()) |v| {
            var val: Device = undefined;
            @memcpy(&val, v[0..3]);
            try vals.append(alloc, val);
        }

        try graph.put(key, vals);
    }

    return graph;
}

fn findPaths(alloc: std.mem.Allocator, graph: *const std.AutoHashMap([3]u8, ArrayList([3]u8))) !u64 {
    var paths = std.AutoHashMap(Device, u16).init(alloc);
    defer paths.deinit();
    var queue = try Queue(Device).init(alloc, 6);
    defer queue.deinit();

    try queue.push("you".*);
    try paths.put("you".*, 1);

    while (!queue.empty()) {
        const n = queue.len;
        for (0..n) |_| {
            const d = queue.pop().?;
            const p = paths.get(d).?;
            if (p == 0) continue;
            const neighbors = graph.get(d);
            if (neighbors == null) continue;
            for (neighbors.?.items) |nb| {
                try queue.push(nb);
                try paths.put(nb, p + (paths.get(nb) orelse 0));
            }
            try paths.put(d, 0);
        }
    }

    return paths.get("out".*).?;
}

fn task1(alloc: std.mem.Allocator, input: []const u8) !u64 {
    var graph = try parseGraph(alloc, input);
    defer {
        var iter = graph.valueIterator();
        while (iter.next()) |v| {
            v.deinit(alloc);
        }
        graph.deinit();
    }

    return try findPaths(alloc, &graph);
}

const Device2 = packed struct {
    a: u8,
    b: u8,
    c: u8,
};

fn parseGraph2(alloc: std.mem.Allocator, input: []const u8) !std.AutoHashMap(Device2, ArrayList(Device2)) {
    var line_iter = std.mem.tokenizeScalar(u8, input, '\n');
    var graph = std.AutoHashMap(Device2, ArrayList(Device2)).init(alloc);

    while (line_iter.next()) |line| {
        var iter = std.mem.tokenizeScalar(u8, line, ' ');
        const key_slice = iter.next().?[0..3];
        const key = Device2{ .a = key_slice[0], .b = key_slice[1], .c = key_slice[2] };

        var vals = ArrayList(Device2){};
        while (iter.next()) |v| {
            const val = Device2{ .a = v[0], .b = v[1], .c = v[2] };
            try vals.append(alloc, val);
        }

        try graph.put(key, vals);
    }

    return graph;
}

const PathNode = struct {
    count: u64,
    has_dac: u2,
    has_fft: u2,
};

fn task2(alloc: std.mem.Allocator, input: []const u8) !u64 {
    // Parse input into graph representation
    var graph = try parseGraph2(alloc, input);
    defer {
        var iter = graph.valueIterator();
        while (iter.next()) |v| {
            v.deinit(alloc);
        }
        graph.deinit();
    }

    // Initialize BFS queue and path tracking with special node states
    var queue = try Queue(Device2).init(alloc, 6);
    defer queue.deinit();
    var paths = std.AutoHashMap(Device2, PathNode).init(alloc);
    defer paths.deinit();

    // Define special nodes: start, required intermediate nodes (fft, dac), and end
    const start = Device2{ .a = 's', .b = 'v', .c = 'r' };
    const fft = Device2{ .a = 'f', .b = 'f', .c = 't' };
    const dac = Device2{ .a = 'd', .b = 'a', .c = 'c' };
    const end = Device2{ .a = 'o', .b = 'u', .c = 't' };
    try queue.push(start);
    try paths.put(start, PathNode{ .count = 1, .has_dac = 0, .has_fft = 0 });

    // BFS to find paths that visit both required intermediate nodes
    while (!queue.empty()) {
        const n = queue.len;
        // Process all nodes at current BFS level
        for (0..n) |_| {
            const node = queue.pop().?;
            var path = paths.get(node).?;
            if (path.count == 0) continue; // Skip already processed nodes
            // std.debug.print("{c}{c}{c}\n", .{node.a, node.b, node.c});
            // std.debug.print("{any}\n\n", .{path});

            // Mark if we've reached required intermediate nodes
            if (node == dac) path.has_dac = 1;
            if (node == fft) path.has_fft = 1;

            const neighbors = graph.get(node);
            if (neighbors == null) continue;

            // Propagate path information to all neighbors
            for (neighbors.?.items) |nb| {
                try queue.push(nb);
                const old = paths.get(nb) orelse PathNode{ .count = 0, .has_dac = 0, .has_fft = 0 };
                var val: PathNode = undefined;

                // Prioritize paths that have visited more required nodes
                if (old.has_dac + old.has_fft > path.has_dac + path.has_fft)
                    val = old
                else if (old.has_dac + old.has_fft < path.has_dac + path.has_fft)
                    val = path
                else
                    // If equal progress, combine path counts and preserve visited flags
                    val = PathNode{ .count = path.count + old.count, .has_dac = @max(path.has_dac, old.has_dac), .has_fft = @max(path.has_fft, old.has_fft) };
                try paths.put(nb, val);
            }
            // Mark current node as processed by zeroing its count
            try paths.put(node, PathNode{ .count = 0, .has_dac = 0, .has_fft = 0 });
        }
    }

    return paths.get(end).?.count;
}

test "task1" {
    const input = comptime @embedFile("inputs/day11_sample.txt");
    const result = try task1(std.testing.allocator, input);
    try std.testing.expectEqual(result, 5);
}

test "task2" {
    const input = comptime @embedFile("inputs/day11_sample2.txt");
    const result = try task2(std.testing.allocator, input);
    try std.testing.expectEqual(result, 2);
}
