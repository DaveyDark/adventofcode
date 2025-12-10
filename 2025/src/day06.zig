const std = @import("std");
const ArrayList = std.ArrayList;

pub fn run(alloc: std.mem.Allocator) !void {
    std.debug.print("Day 6\n", .{});
    const part1 = try task1(alloc, comptime @embedFile("inputs/day6_input.txt"));
    std.debug.print("Part 1: {d}\n", .{part1});
    const part2 = try task2(alloc, comptime @embedFile("inputs/day6_input.txt"));
    std.debug.print("Part 2: {d}\n", .{part2});
}

fn task1(allocator: std.mem.Allocator, input: []const u8) !u64 {
    var line_iter = std.mem.tokenizeScalar(u8, input, '\n');

    var buf = ArrayList(ArrayList(u16)){};
    defer {
        for (buf.items) |*r| {
            r.deinit(allocator);
        }
        buf.deinit(allocator);
    }
    var operators = ArrayList(u8){};
    defer operators.deinit(allocator);

    while (line_iter.next()) |line| {
        var num_iter = std.mem.tokenizeScalar(u8, line, ' ');
        var row_list = ArrayList(u16){};
        while (num_iter.next()) |num| {
            if (num[0] == '*' or num[0] == '+') {
                try operators.append(allocator, num[0]);
                continue;
            }
            const n = try std.fmt.parseInt(u16, num, 10);
            try row_list.append(allocator, n);
        }
        if (row_list.items.len == 0) break;
        try buf.append(allocator, row_list);
    }

    var res: u64 = 0;

    for (operators.items, 0..operators.items.len) |op, i| {
        var col_res: u64 = if (op == '+') 0 else 1;
        for (0..buf.items.len) |j| {
            if (op == '+') {
                col_res += buf.items[j].items[i];
            } else {
                col_res *= buf.items[j].items[i];
            }
        }
        res += col_res;
    }

    return res;
}

pub fn transposeText(
    allocator: std.mem.Allocator,
    input: []const u8,
) ![]u8 {
    // --- Split input into lines (tokens, not preserving empty lines) ---
    var lines = std.mem.tokenizeAny(u8, input, "\n");

    var rows = std.ArrayList([]const u8){};
    defer rows.deinit(allocator);

    while (lines.next()) |line| {
        try rows.append(allocator, line);
    }

    // --- Compute dimensions ---
    var max_len: usize = 0;
    for (rows.items) |r| {
        max_len = @max(max_len, r.len);
    }

    const row_count = rows.items.len;
    const col_count = max_len;

    // Allocate worst-case output (fully padded)
    const out_len = col_count * (row_count + 1);
    var out = try allocator.alloc(u8, out_len);

    var index: usize = 0;

    // --- Transpose ---
    for (0..col_count) |col| {
        for (0..row_count) |row| {
            const line = rows.items[row];
            out[index] = if (col < line.len) line[col] else ' ';
            index += 1;
        }
        out[index] = '\n';
        index += 1;
    }

    return out[0..index];
}

fn parseTransposedNumbers(
    allocator: std.mem.Allocator,
    chunks: []const u8,
) !ArrayList(ArrayList(u16)) {
    var nums = ArrayList(ArrayList(u16)){};

    var iter = std.mem.tokenizeScalar(u8, chunks, '\n');
    var curr = ArrayList(u16){};

    while (iter.next()) |chunk| {
        const trimmed = std.mem.trim(u8, chunk, " ");

        if (trimmed.len == 0) {
            try nums.append(allocator, curr);
            curr = ArrayList(u16){};
            continue;
        }

        const n = try std.fmt.parseInt(u16, trimmed, 10);
        try curr.append(allocator, n);
    }

    try nums.append(allocator, curr);
    return nums;
}

fn parseOps(allocator: std.mem.Allocator, input: []const u8) !ArrayList(u8) {
    var ops = ArrayList(u8){};
    var it = std.mem.tokenizeScalar(u8, input, ' ');

    while (it.next()) |tok| {
        try ops.append(allocator, tok[0]);
    }

    return ops;
}

fn task2(
    allocator: std.mem.Allocator,
    input: []const u8,
) !u64 {
    // --- Split input into number block and ops block ---
    const split_idx = std.mem.lastIndexOf(u8, input, "\n").?;
    const nums_str = input[0..split_idx];
    const ops_str = input[split_idx + 1 ..];

    // --- Transpose the number grid ---
    const chunks = try transposeText(allocator, nums_str);
    defer allocator.free(chunks);

    // --- Parse numeric columns ---
    var nums = try parseTransposedNumbers(allocator, chunks);
    defer {
        for (nums.items) |*arr| arr.deinit(allocator);
        nums.deinit(allocator);
    }

    // --- Parse operators ---
    var ops = try parseOps(allocator, ops_str);
    defer ops.deinit(allocator);

    // --- Solve each column ---
    var res: u64 = 0;

    for (ops.items, 0..) |op, i| {
        const data = nums.items[i].items;

        var subtotal: u64 = if (op == '*') 1 else 0;

        for (data) |n| {
            if (op == '*') subtotal *= n else subtotal += n;
        }

        res += subtotal;
    }

    return res;
}

test "task1" {
    const input = comptime @embedFile("inputs/day6_sample.txt");
    const result = try task1(std.testing.allocator, input);
    try std.testing.expectEqual(result, 4277556);
}

test "task2" {
    const input = comptime @embedFile("inputs/day6_sample.txt");
    const result = try task2(std.testing.allocator, input);
    try std.testing.expectEqual(result, 3263827);
}
