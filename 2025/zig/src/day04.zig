const std = @import("std");
const ArrayList = std.ArrayList;
const Grid = @import("common/grid.zig").Grid;

pub fn run(alloc: std.mem.Allocator) !void {
    std.debug.print("Day 4\n", .{});
    const part1 = try task1(alloc, comptime @embedFile("inputs/day4_input.txt"));
    std.debug.print("Part 1: {d}\n", .{part1});
    const part2 = try task2(alloc, comptime @embedFile("inputs/day4_input.txt"));
    std.debug.print("Part 2: {d}\n", .{part2});
}

const GridTile = enum { empty, roll };

fn conv_fn(c: u8) GridTile {
    return switch (c) {
        '.' => GridTile.empty,
        '@' => GridTile.roll,
        else => unreachable,
    };
}

fn fmt_fn(c: GridTile) void {
    const ch: u8 = switch (c) {
        GridTile.empty => ' ',
        GridTile.roll => '@',
    };
    std.debug.print("{c}", .{ch});
}

fn isAccessible(grid: Grid(GridTile), i: usize, j: usize) bool {
    const adj = [8][2]i8{
        .{ -1, -1 },
        .{ -1, 0 },
        .{ -1, 1 },
        .{ 0, -1 },
        .{ 0, 1 },
        .{ 1, -1 },
        .{ 1, 0 },
        .{ 1, 1 },
    };

    var paper: u32 = 0;
    for (adj) |dir| {
        const di = @as(isize, @intCast(i)) + dir[0];
        const dj = @as(isize, @intCast(j)) + dir[1];
        if (grid.isOob(di, dj)) continue;
        if (grid.get(@intCast(di), @intCast(dj)) == GridTile.empty) continue;
        paper += 1;
    }
    return (paper < 4);
}

fn task1(allocator: std.mem.Allocator, input: []const u8) !u64 {
    var grid = try Grid(GridTile).newFromInput(allocator, input, conv_fn);
    defer grid.deinit(allocator);

    var accessible_rolls: u32 = 0;

    for (0..grid.rows) |i| {
        for (0..grid.cols) |j| {
            if (grid.get(i, j) == GridTile.empty) continue;
            if (isAccessible(grid, i, j)) {
                accessible_rolls += 1;
            }
        }
    }

    return accessible_rolls;
}

fn task2(allocator: std.mem.Allocator, input: []const u8) !u64 {
    var grid = try Grid(GridTile).newFromInput(allocator, input, conv_fn);
    defer grid.deinit(allocator);

    var accessible_rolls: u32 = 0;
    var last: ?u32 = null;
    while (last == null or last != accessible_rolls) {
        last = accessible_rolls;
        for (0..grid.rows) |i| {
            for (0..grid.cols) |j| {
                if (grid.get(i, j) == GridTile.empty) continue;
                if (isAccessible(grid, i, j)) {
                    grid.set(i, j, GridTile.empty);
                    accessible_rolls += 1;
                }
            }
        }
    }

    return accessible_rolls;
}

test "task1" {
    const input = comptime @embedFile("inputs/day4_sample.txt");
    const result = try task1(std.testing.allocator, input);
    try std.testing.expectEqual(result, 13);
}

test "task2" {
    const input = comptime @embedFile("inputs/day4_sample.txt");
    const result = try task2(std.testing.allocator, input);
    try std.testing.expectEqual(result, 43);
}
