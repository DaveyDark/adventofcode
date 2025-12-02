const std = @import("std");
const ArrayList = std.ArrayList;

// Day 1: Secret Enterance

pub fn run(alloc: std.mem.Allocator) !void {
    std.debug.print("Day 1\n", .{});
    const part1 = try task1(alloc, comptime @embedFile("inputs/day1_input.txt"));
    std.debug.print("Part 1: {d}\n", .{part1});
    const part2 = try task2(alloc, comptime @embedFile("inputs/day1_input.txt"));
    std.debug.print("Part 2: {d}\n", .{part2});
}

// Enum to represent the direction of dial movement
const Direction = enum {
    left,   
    right, 
};

// Part 1: Count how many times the dial lands exactly on position 0
// Approach: Start at position 50, apply each movement, and count zeros
fn task1(_: std.mem.Allocator, input: []const u8) !u32 {
    var lineIter = std.mem.tokenizeAny(u8, input, "\n");
    var dial: i32 = 50;  // Starting position on the 0-99 dial
    var count: u32 = 0;  // Count of times we land on 0

    // Process each line of input (format: "L5" or "R3" for left/right + magnitude)
    while (lineIter.next()) |line| {
        if (line.len == 0) break;
        
        // Parse direction 
        const direction: Direction = if (line[0] == 'L') .left else .right;
        // Parse magnitude
        const magnitude = try std.fmt.parseInt(i32, line[1..], 10);

        // Apply the movement to the dial 
        switch (direction) {
            .left => dial = @mod(dial - magnitude, 100),   // Move counter-clockwise
            .right => dial = @mod(dial + magnitude, 100),  // Move clockwise
        }

        // Check if we landed exactly on 0
        if (dial == 0) count += 1;
    }

    return count;
}

// Part 2: Count how many times the dial passes through position 0 during movement
// Approach: Calculate how many complete rotations occur during each move
fn task2(_: std.mem.Allocator, input: []const u8) !u32 {
    var lineIter = std.mem.tokenizeAny(u8, input, "\n");
    var dial: i32 = 50;  // Starting position on the 0-99 dial
    var count: u32 = 0;  // Count of times we pass through 0

    // Process each movement command
    while (lineIter.next()) |line| {
        if (line.len == 0) break;
        
        // Parse direction and magnitude
        const direction: Direction = if (line[0] == 'L') .left else .right;
        const magnitude = try std.fmt.parseInt(i32, line[1..], 10);
        const isZero = dial == 0;  // Remember if we started at 0

        switch (direction) {
            .left => {
                // Moving left (counter-clockwise)
                dial -= magnitude;
                
                // If we went past 0 (into negative), we passed through 0
                if (dial <= 0) {
                    // Count complete rotations (every 100 units = 1 pass through 0)
                    count += @intCast(@divFloor(-dial, 100));
                    // Add 1 more pass if we didn't start at 0 (first crossing)
                    count += if (isZero) 0 else 1;
                }
                // Normalize dial position back to 0-99 range
                dial = @mod(dial, 100);
            },
            .right => {
                // Moving right (clockwise)
                dial += magnitude;
                
                // If we went past 99, we passed through 0
                if (dial >= 100) {
                    // Count how many complete rotations we made
                    count += @intCast(@divFloor(dial, 100));
                }
                // Normalize dial position back to 0-99 range
                dial = @mod(dial, 100);
            },
        }
    }

    return count;
}

test "task1" {
    const input = comptime @embedFile("inputs/day1_sample.txt");
    const result = try task1(std.testing.allocator, input);
    try std.testing.expectEqual(result, 3);
}

test "task2" {
    const input = comptime @embedFile("inputs/day1_sample.txt");
    const result = try task2(std.testing.allocator, input);
    try std.testing.expectEqual(result, 6);
}
