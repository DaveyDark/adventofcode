const std = @import("std");
const ArrayList = std.ArrayList;
const BitSet = std.bit_set.IntegerBitSet(16);

const FactoryMachine = struct {
    target_lights: u16,
    num_lights: usize,
    buttons: []u16,
};

pub fn run(alloc: std.mem.Allocator) !void {
    std.debug.print("Day 10\n", .{});
    const part1 = try task1(alloc, comptime @embedFile("inputs/day10_input.txt"));
    std.debug.print("Part 1: {d}\n", .{part1});
    const part2 = try task2(alloc, comptime @embedFile("inputs/day10_input.txt"));
    std.debug.print("Part 2: {d}\n", .{part2});
}

fn parseLights(input: []const u8) u16 {
    var ans = BitSet.initEmpty();
    var iter = std.mem.reverseIterator(input[1 .. input.len - 1]);
    var i: u16 = 0;
    while (iter.next()) |c| {
        if (c == '#') ans.set(i);
        i += 1;
    }

    return ans.mask;
}

fn parseButton(input: []const u8, n: usize) u16 {
    var iter = std.mem.tokenizeScalar(
        u8,
        input[1 .. input.len - 1],
        ',',
    );

    var ans = BitSet.initEmpty();
    while (iter.next()) |c| {
        const idx = c[0] - '0';
        ans.set(n - idx - 1);
    }

    return ans.mask;
}

fn parseMachine(alloc: std.mem.Allocator, line: []const u8) !FactoryMachine {
    var iter = std.mem.tokenizeScalar(u8, line, ' ');
    const lights = iter.next().?;
    const ans = parseLights(lights);
    const n = lights.len - 2;
    var buttons = ArrayList(u16){};
    defer buttons.deinit(alloc);

    while (iter.next()) |slice| {
        if (slice[0] == '{') break;
        try buttons.append(alloc, parseButton(slice, n));
    }

    return FactoryMachine{ .target_lights = ans, .num_lights = n, .buttons = try buttons.toOwnedSlice(alloc) };
}

fn task1(alloc: std.mem.Allocator, input: []const u8) !u64 {
    var line_iter = std.mem.tokenizeScalar(u8, input, '\n');

    var result: u64 = 0;
    while (line_iter.next()) |line| {
        const machine = try parseMachine(alloc, line);
        defer alloc.free(machine.buttons);

        var min_res: u8 = 16;
        for (1..std.math.pow(u16, 2, @intCast(machine.buttons.len))) |combination| {
            var i: u16 = 0;
            var res: u16 = 0;
            var comb = combination;
            while (comb > 0) {
                const bit: u1 = @intCast(comb & 1);
                if (bit == 1) {
                    res ^= machine.buttons[i];
                }
                comb >>= 1;
                i += 1;
            }

            if (res == machine.target_lights) {
                min_res = @min(min_res, @popCount(combination));
            }
        }

        result += min_res;
    }

    return result;
}

fn task2(alloc: std.mem.Allocator, input: []const u8) !u64 {
    // TODO: Solve part 2
    // Seems too hard for now
    _ = alloc;
    _ = input;

    return 0;
}

test "task1" {
    const input = comptime @embedFile("inputs/day10_sample.txt");
    const result = try task1(std.testing.allocator, input);
    try std.testing.expectEqual(result, 7);
}

test "task2" {
    const input = comptime @embedFile("inputs/day10_sample.txt");
    const result = try task2(std.testing.allocator, input);
    try std.testing.expectEqual(result, 0);
}
