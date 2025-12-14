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

fn task1(alloc: std.mem.Allocator, input: []const u8) !u64 {
    var line_iter = std.mem.tokenizeScalar(u8, input, '\n');

    var graph = std.AutoHashMap([3]u8, ArrayList([3]u8)).init(alloc);
    defer {
        var iter = graph.valueIterator();
        while (iter.next()) |v| {
            v.deinit(alloc);
        }
        graph.deinit();
    }

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

fn task2(alloc: std.mem.Allocator, input: []const u8) !u64 {
    _ = alloc;
    _ = input;

    return 0;
}

test "task1" {
    const input = comptime @embedFile("inputs/day11_sample.txt");
    const result = try task1(std.testing.allocator, input);
    try std.testing.expectEqual(result, 5);
}

test "task2" {
    const input = comptime @embedFile("inputs/day11_sample.txt");
    const result = try task2(std.testing.allocator, input);
    try std.testing.expectEqual(result, 0);
}
