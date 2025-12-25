const std = @import("std");
const ArrayList = std.ArrayList;

pub fn run(alloc: std.mem.Allocator) !void {
    std.debug.print("Day 7\n", .{});
    const part1 = try task1(alloc, comptime @embedFile("inputs/day7_input.txt"));
    std.debug.print("Part 1: {d}\n", .{part1});
    const part2 = try task2(alloc, comptime @embedFile("inputs/day7_input.txt"));
    std.debug.print("Part 2: {d}\n", .{part2});
}

fn task1(alloc: std.mem.Allocator, input: []const u8) !u32 {
    const split_idx = std.mem.indexOfScalar(u8, input, '\n').?;
    const header = input[0..split_idx];
    const lines = input[split_idx + 1 ..];
    var line_iter = std.mem.tokenizeScalar(u8, lines, '\n');

    var beam = ArrayList(u1){};
    defer beam.deinit(alloc);
    for (header) |ch| {
        if (ch == 'S') try beam.append(alloc, 1) else try beam.append(alloc, 0);
    }

    var splits: u32 = 0;

    while (line_iter.next()) |line| {
        for (line, 0..line.len) |ch, i| {
            if (ch == '^' and beam.items[i] == 1) {
                beam.items[i] = 0;
                if (i > 0) beam.items[i - 1] = 1;
                if (i + 1 != beam.items.len) beam.items[i + 1] = 1;
                splits += 1;
            }
        }
    }

    return splits;
}

fn task2(alloc: std.mem.Allocator, input: []const u8) !u64 {
    const split_idx = std.mem.indexOfScalar(u8, input, '\n').?;
    const header = input[0..split_idx];
    const lines = input[split_idx + 1 ..];
    var line_iter = std.mem.tokenizeScalar(u8, lines, '\n');

    var beam = ArrayList(u64){};
    defer beam.deinit(alloc);
    for (header) |ch| {
        if (ch == 'S') try beam.append(alloc, 1) else try beam.append(alloc, 0);
    }

    while (line_iter.next()) |line| {
        for (line, 0..line.len) |ch, i| {
            if (ch == '^' and beam.items[i] != 0) {
                const power = beam.items[i];
                beam.items[i] = 0;
                if (i > 0) beam.items[i - 1] += power;
                if (i + 1 != beam.items.len) beam.items[i + 1] += power;
            }
        }
    }

    var sum: u64 = 0;
    for (beam.items) |b| {
        sum += b;
    }

    return sum;
}

test "task1" {
    const input = comptime @embedFile("inputs/day7_sample.txt");
    const result = try task1(std.testing.allocator, input);
    try std.testing.expectEqual(result, 21);
}

test "task2" {
    const input = comptime @embedFile("inputs/day7_sample.txt");
    const result = try task2(std.testing.allocator, input);
    try std.testing.expectEqual(result, 40);
}
