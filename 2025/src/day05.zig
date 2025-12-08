const std = @import("std");
const ArrayList = std.ArrayList;
const RangeGroup = @import("common/range.zig").RangeGroup;
const utils = @import("common/utils.zig");

pub fn run(alloc: std.mem.Allocator) !void {
    std.debug.print("Day 5\n", .{});
    const part1 = try task1(alloc, comptime @embedFile("inputs/day5_input.txt"));
    std.debug.print("Part 1: {d}\n", .{part1});
    const part2 = try task2(alloc, comptime @embedFile("inputs/day5_input.txt"));
    std.debug.print("Part 2: {d}\n", .{part2});
}

fn task1(allocator: std.mem.Allocator, input: []const u8) !u32 {
    var inputIter = std.mem.splitScalar(u8, input, '\n');
    var ranges = RangeGroup(u64).newRange();
    defer ranges.deinit(allocator);

    while (inputIter.next()) |line| {
        if (utils.isStringEmpty(line)) break;
        var splitIter = std.mem.tokenizeAny(u8, line, "-");
        const start = try std.fmt.parseInt(u64, splitIter.next().?, 10);
        const end = try std.fmt.parseInt(u64, splitIter.next().?, 10);

        try ranges.addRange(allocator, [2]u64{ start, end });
    }

    var valid: u32 = 0;
    while (inputIter.next()) |line| {
        const val = try std.fmt.parseInt(u64, line, 10);

        if (ranges.containsValue(val)) valid += 1;
    }

    return valid;
}

fn task2(allocator: std.mem.Allocator, input: []const u8) !u64 {
    var inputIter = std.mem.splitScalar(u8, input, '\n');
    var ranges = RangeGroup(u64).newRange();
    defer ranges.deinit(allocator);

    while (inputIter.next()) |line| {
        if (utils.isStringEmpty(line)) break;
        var splitIter = std.mem.tokenizeAny(u8, line, "-");
        const start = try std.fmt.parseInt(u64, splitIter.next().?, 10);
        const end = try std.fmt.parseInt(u64, splitIter.next().?, 10);

        try ranges.addRange(allocator, [2]u64{ start, end });
    }

    var total: u64 = 0;
    for (ranges.ranges.items) |rng| {
        total += rng[1] - rng[0] + 1;
    }

    return total;
}

test "task1" {
    const input = comptime @embedFile("inputs/day5_sample.txt");
    const result = try task1(std.testing.allocator, input);
    try std.testing.expectEqual(result, 3);
}

test "task2" {
    const input = comptime @embedFile("inputs/day5_sample.txt");
    const result = try task2(std.testing.allocator, input);
    try std.testing.expectEqual(result, 14);
}
