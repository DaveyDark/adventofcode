const std = @import("std");

pub fn run(alloc: std.mem.Allocator) !void {
    std.debug.print("Day 2\n", .{});
    const part1 = try task1(alloc, comptime @embedFile("inputs/day3_input.txt"));
    std.debug.print("Part 1: {d}\n", .{part1});
    const part2 = try task2(alloc, comptime @embedFile("inputs/day3_input.txt"));
    std.debug.print("Part 2: {d}\n", .{part2});
}

fn processBank(alloc: std.mem.Allocator, bank: []const u8, k: usize) !u64 {
    const buf = try alloc.alloc(u8, k); // Use as stack
    var buf_top: usize = 0;
    defer alloc.free(buf);

    var remaining_chars: usize = bank.len;

    for (bank) |c| {
        remaining_chars -= 1;

        if (buf_top == k and buf[buf_top - 1] >= c) {
            // No need to push
            continue;
        }

        // Pop until top element is >= c, but ensure we can still fill the buffer
        while (buf_top > 0 and buf[buf_top - 1] < c and (buf_top + remaining_chars >= k)) {
            buf_top -= 1;
        }

        // push element
        buf[buf_top] = c;
        buf_top += 1;
    }

    const ans = try std.fmt.parseInt(u64, buf, 10);

    return ans;
}

fn task1(allocator: std.mem.Allocator, input: []const u8) !u64 {
    var lineIter = std.mem.tokenizeAny(u8, input, "\n");
    var ans: u64 = 0;

    while (lineIter.next()) |line| {
        if (line.len == 0) break;
        ans += try processBank(allocator, line, 2);
    }

    return ans;
}

fn task2(allocator: std.mem.Allocator, input: []const u8) !u64 {
    var lineIter = std.mem.tokenizeAny(u8, input, "\n");
    var ans: u64 = 0;

    while (lineIter.next()) |line| {
        if (line.len == 0) break;
        ans += try processBank(allocator, line, 12);
    }

    return ans;
}

test "task1" {
    const input = comptime @embedFile("inputs/day3_sample.txt");
    const result = try task1(std.testing.allocator, input);
    try std.testing.expectEqual(result, 357);
}

test "task1-edge" {
    // When input involves using the last digit
    const input =
        \\8900009
    ;
    const result = try task1(std.testing.allocator, input);
    try std.testing.expectEqual(result, 99);
}

test "task1-edge2" {
    // When input is exactly 2 digits
    const input =
        \\89
    ;
    const result = try task1(std.testing.allocator, input);
    try std.testing.expectEqual(result, 89);
}

test "task1-edge3" {
    const input =
        \\190004
    ;
    const result = try task1(std.testing.allocator, input);
    try std.testing.expectEqual(result, 94);
}

test "task1-edge4" {
    const input =
        \\619
    ;
    const result = try task1(std.testing.allocator, input);
    try std.testing.expectEqual(result, 69);
}

test "task1-edge5" {
    const input =
        \\997
    ;
    const result = try task1(std.testing.allocator, input);
    try std.testing.expectEqual(result, 99);
}

test "task2" {
    const input = comptime @embedFile("inputs/day3_sample.txt");
    const result = try task2(std.testing.allocator, input);
    try std.testing.expectEqual(result, 3121910778619);
}
