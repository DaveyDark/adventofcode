const std = @import("std");
const ArrayList = std.ArrayList;
const Grid = @import("common/grid.zig").Grid(GridTile);

const GridTile = enum { White, Red };

fn convFn(c: u8) GridTile {
    return switch (c) {
        '.' => GridTile.White,
        '#' => GridTile.Red,
        else => unreachable,
    };
}

pub fn run(alloc: std.mem.Allocator) !void {
    std.debug.print("Day 9\n", .{});
    const part1 = try task1(alloc, comptime @embedFile("inputs/day9_input.txt"));
    std.debug.print("Part 1: {d}\n", .{part1});
    const part2 = try task2(alloc, comptime @embedFile("inputs/day9_input.txt"));
    std.debug.print("Part 2: {d}\n", .{part2});
}

inline fn idx(x: usize, y: usize, n: usize) usize {
    return x * n + y;
}

inline fn calculateArea(arr: ArrayList([2]usize), i: usize, j: usize) usize {
    const p1 = arr.items[i];
    const p2 = arr.items[j];
    return (@max(p1[0], p2[0]) - @min(p1[0], p2[0]) + 1) * (@max(p1[1], p2[1]) - @min(p1[1], p2[1]) + 1);
}

fn task1(alloc: std.mem.Allocator, input: []const u8) !u64 {
    var line_iter = std.mem.tokenizeScalar(u8, input, '\n');
    var tiles = ArrayList([2]usize){};
    defer tiles.deinit(alloc);

    while (line_iter.next()) |line| {
        var num_iter = std.mem.tokenizeScalar(u8, line, ',');
        const i = try std.fmt.parseInt(usize, num_iter.next().?, 10);
        const j = try std.fmt.parseInt(usize, num_iter.next().?, 10);
        try tiles.append(alloc, [2]usize{ i, j });
    }

    const n = tiles.items.len;
    var max_area: u64 = 0;

    for (0..n) |i| {
        for (i+1..n) |j| {
            max_area = @max(max_area, calculateArea(tiles, i, j));
        }
    }

    return max_area;
}

fn task2(alloc: std.mem.Allocator, input: []const u8) !u64 {
    _ = alloc;
    _ = input;
    return 0;
}

test "task1" {
    const input = comptime @embedFile("inputs/day9_sample.txt");
    const result = try task1(std.testing.allocator, input);
    try std.testing.expectEqual(result, 50);
}

test "task2" {
    const input = comptime @embedFile("inputs/day9_sample.txt");
    const result = try task2(std.testing.allocator, input);
    try std.testing.expectEqual(result, 0);
}
