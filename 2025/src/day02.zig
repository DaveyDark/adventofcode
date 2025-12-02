const std = @import("std");
const ArrayList = std.ArrayList;

// Day 2: Gift Shop
pub fn run(alloc: std.mem.Allocator) !void {
    std.debug.print("Day 2\n", .{});
    const part1 = try task1(alloc, comptime @embedFile("inputs/day2_input.txt"));
    std.debug.print("Part 1: {d}\n", .{part1});
    const part2 = try task2(alloc, comptime @embedFile("inputs/day2_input.txt"));
    std.debug.print("Part 2: {d}\n", .{part2});
}

// For Part 1, Check if the ID is invalid
// If a string is a substring repeated twice, it is invalid
fn is_invalid(id: u64, str_buf: *[10]u8) !bool {
    var size: usize = 0;

    // Format the ID into a string
    var num = id;
    while (num > 0) : (num /= 10) {
        const c = num % 10;
        str_buf[size] = '0' + @as(u8, @intCast(c));
        size += 1;
    }

    // If Odd length string, ignore
    if (size % 2 != 0) return false;

    // Checl if left and right halves are equal
    return std.mem.eql(u8, str_buf[0..size/2], str_buf[size/2..size]);
}

// Part 1: Sum all invalid IDs within given ranges
fn task1(_: std.mem.Allocator, input: []const u8) !u64 {
    var rangeIter = std.mem.tokenizeAny(u8, input, ",");
    var ans: u64 = 0;
    var str_buf: [10]u8 = undefined;

    // Parse comma-separated ranges from input (format: "start1-end1,start2-end2,...")
    while(rangeIter.next()) |range| {
        // Split each range on hyphen to get start and end values
        var splitIter = std.mem.tokenizeAny(u8, range, "-");
        const start = try std.fmt.parseInt(u64, splitIter.next().?, 10);
        const end = try std.fmt.parseInt(u64, splitIter.next().?, 10);

        // Check every ID in the range
        var n = start;
        while (n <= end) : (n += 1) {
            // If the ID is invalid (meets our criteria), add it to the sum
            if(try is_invalid(n, &str_buf)) {
                ans += n;
            }
        }
    }

    return ans;
}

// For Part 2, Check if the ID is invalid
// If a string is a substring repeated ANY NUMBER OF TIMES, it is invalid
fn is_invalid2(id: u64, str_buf: *[10]u8) !bool {
    var size: usize = 0;

    // Convert ID to string representation and store in source buffer
    var num = id;
    while (num > 0) : (num /= 10) {
        const c = num % 10;
        str_buf[size] = '0' + @as(u8, @intCast(c));
        size += 1;
    }

    // Size of left and right substrings
    var left_length = size - 1;
    var right_length: usize = 1;

    while (left_length >= right_length) {
        // If substrings are not divisible
        if(left_length % right_length != 0) {
            right_length += 1;
            left_length = size - right_length;
            continue;
        }
        // Calculate number of repetitions
        const repeats = @max(left_length / right_length, 1);

        // Check if substring is repeated expected number of times
        if(std.mem.containsAtLeast(u8, str_buf[0..left_length], repeats, str_buf[left_length..size])) {
            return true;
        }


        // Update right_length and left_length
        right_length += 1;
        left_length = size - right_length;
    }
    return false;
}

// Part 2: Currently not implemented - returns 0
// Would need to process input and check IDs using is_invalid2
fn task2(_: std.mem.Allocator, input: []const u8) !u64 {
    var rangeIter = std.mem.tokenizeAny(u8, input, ",");
    var ans: u64 = 0;
    var str_buf: [10]u8 = undefined;

    // Parse ranges from input
    while(rangeIter.next()) |range| {
        var splitIter = std.mem.tokenizeAny(u8, range, "-");
        const start = try std.fmt.parseInt(u64, splitIter.next().?, 10);
        const end = try std.fmt.parseInt(u64, splitIter.next().?, 10);

        // Loop over the range of IDs
        var n = start;
        while (n <= end) : (n += 1) {
            if(try is_invalid2(n, &str_buf)) {
                ans += n;
            }
        }
    }

    return ans;
}

test "task1" {
    const input = comptime @embedFile("inputs/day2_sample.txt");
    const result = try task1(std.testing.allocator, input);
    try std.testing.expectEqual(result, 1227775554);
}

test "task2" {
    const input = comptime @embedFile("inputs/day2_sample.txt");
    const result = try task2(std.testing.allocator, input);
    try std.testing.expectEqual(result, 4174379265);
}
