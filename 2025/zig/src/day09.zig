const std = @import("std");
const ArrayList = std.ArrayList;
const Grid = @import("common/grid.zig").Grid(GridTile);

const GridTile = enum { White, Green, Red, Unset };

fn fmt_fn(c: GridTile) void {
    const ch: u8 = switch (c) {
        GridTile.Unset => '.',
        GridTile.White => ' ',
        GridTile.Red => '#',
        GridTile.Green => 'X',
    };
    std.debug.print("{c}", .{ch});
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
        for (i + 1..n) |j| {
            max_area = @max(max_area, calculateArea(tiles, i, j));
        }
    }

    return max_area;
}

fn connectTiles(grid: *Grid, p1: [2]usize, p2: [2]usize) void {
    if (p1[0] == p2[0]) {
        const lower = @min(p1[1], p2[1]);
        const upper = @max(p1[1], p2[1]);
        for (lower + 1..upper) |j| {
            grid.set(p1[0], j, GridTile.Green);
        }
    } else if (p1[1] == p2[1]) {
        const lower = @min(p1[0], p2[0]);
        const upper = @max(p1[0], p2[0]);
        for (lower + 1..upper) |i| {
            grid.set(i, p1[1], GridTile.Green);
        }
    }
}

fn floodFill(alloc: std.mem.Allocator, grid: *Grid) !void {
    var stack = ArrayList([2]usize){};
    defer stack.deinit(alloc);
    try stack.append(alloc, [2]usize{ 0, 0 });

    while (stack.items.len != 0) {
        const p = stack.pop().?;
        grid.set(p[0], p[1], GridTile.White);
        if (p[0] > 0 and grid.get(p[0] - 1, p[1]) == GridTile.Unset)
            try stack.append(alloc, [2]usize{ p[0] - 1, p[1] });
        if (p[1] > 0 and grid.get(p[0], p[1] - 1) == GridTile.Unset)
            try stack.append(alloc, [2]usize{ p[0], p[1] - 1 });
        if (p[0] != grid.rows - 1 and grid.get(p[0] + 1, p[1]) == GridTile.Unset)
            try stack.append(alloc, [2]usize{ p[0] + 1, p[1] });
        if (p[1] != grid.cols - 1 and grid.get(p[0], p[1] + 1) == GridTile.Unset)
            try stack.append(alloc, [2]usize{ p[0], p[1] + 1 });
    }
}

fn task2(alloc: std.mem.Allocator, input: []const u8) !u64 {
    var line_iter = std.mem.tokenizeScalar(u8, input, '\n');
    var tiles = ArrayList([2]usize){};
    defer tiles.deinit(alloc);

    // FIX: Too inefficient approach, real input it too big to create a grid from (>9gb of RAM)

    // Min i, Min j, Max i, Max j
    var bounds = [4]usize{ std.math.maxInt(usize), std.math.maxInt(usize), 0, 0 };

    while (line_iter.next()) |line| {
        var num_iter = std.mem.tokenizeScalar(u8, line, ',');
        const j = try std.fmt.parseInt(usize, num_iter.next().?, 10);
        const i = try std.fmt.parseInt(usize, num_iter.next().?, 10);
        bounds[0] = @min(bounds[0], i);
        bounds[1] = @min(bounds[1], j);
        bounds[2] = @max(bounds[2], i);
        bounds[3] = @max(bounds[3], j);
        try tiles.append(alloc, [2]usize{ i, j });
    }

    // Normalize Tiles
    for (tiles.items) |*t| {
        t[0] -= bounds[0] - 1;
        t[1] -= bounds[1] - 1;
    }

    // Build Grid
    const w = bounds[3] - bounds[1] + 3;
    const h = bounds[2] - bounds[0] + 3;

    const items = try alloc.alloc(GridTile, w * h);
    for (items) |*i| {
        i.* = GridTile.Unset;
    }
    var grid = Grid.new(items, h, w);
    defer grid.deinit(alloc);

    for (tiles.items) |t| {
        grid.set(t[0], t[1], GridTile.Red);
    }

    // Add Green Tiles
    var window_iter = std.mem.window([2]usize, tiles.items, 2, 1);
    while (window_iter.next()) |win| {
        connectTiles(&grid, win[0], win[1]);
    }
    connectTiles(&grid, tiles.items[0], tiles.getLast());
    try floodFill(alloc, &grid);

    for (items) |*i| {
        if (i.* == GridTile.Unset)
            i.* = GridTile.Green;
    }

    grid.customPrint(fmt_fn);

    // TODO get largest rect in grid (largest rect in histogram algo)

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
    try std.testing.expectEqual(result, 24);
}
