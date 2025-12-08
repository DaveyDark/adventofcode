const std = @import("std");
const ArrayList = std.ArrayList;
// const Grid = struct {};

pub fn Grid(comptime T: type) type {
    return struct {
        items: []T,
        rows: usize,
        cols: usize,

        pub fn get(self: *const @This(), i: usize, j: usize) T {
            return self.items[i * self.cols + j];
        }

        pub fn set(self: *@This(), i: usize, j: usize, value: T) void {
            self.items[i * self.cols + j] = value;
        }

        pub fn isOob(self: *const @This(), i: isize, j: isize) bool {
            return i < 0 or j < 0 or i >= self.rows or j >= self.cols;
        }

        pub fn new(buf: []T, rows: usize, cols: usize) @This() {
            return @This(){ .items = buf, .rows = rows, .cols = cols };
        }

        pub fn deinit(self: *@This(), allocator: std.mem.Allocator) void {
            allocator.free(self.items);
        }

        pub fn customPrint(self: *const @This(), fmt_fn: fn (value: T) void) void {
            for (0..self.rows) |i| {
                for (0..self.cols) |j| {
                    const v = self.get(i, j);
                    fmt_fn(v);
                }
                std.debug.print("\n", .{});
            }
        }

        pub fn prettyPrint(self: *const @This()) void {
            const info = @typeInfo(T);

            const fmt_fn = struct {
                fn format(v: T) void {
                    switch (info) {
                        .int => |int_info| {
                            if (int_info.signedness == .unsigned and int_info.bits == 8) {
                                // u8 type
                                // print unsigned byte as ASCII char
                                std.debug.print("{c}", .{v});
                            } else {
                                std.debug.print("{d}", .{v});
                            }
                        },
                        else => std.debug.print("{} ", .{v}),
                    }
                }
            }.format;

            self.customPrint(fmt_fn);
        }

        pub fn newFromInput(allocator: std.mem.Allocator, input: []const u8, comptime conv_fn: ?fn (c: u8) T) !@This() {
            var list = ArrayList(T){};
            defer list.deinit(allocator);

            var rows: usize = 1;
            var cols: usize = 0;
            for (input) |ch| {
                cols += 1;
                if (ch == '\n') {
                    rows += 1;
                    cols = 0;
                    continue;
                }
                try list.append(allocator, if (conv_fn) |f| f(ch) else @as(T, ch));
            }

            const buf = try allocator.alloc(T, list.items.len);
            @memcpy(buf, list.items);

            return @This().new(buf, rows, cols);
        }
    };
}

test "Grid basic operations" {
    const allocator = std.testing.allocator;

    // Test with i32 type
    const buf = try allocator.alloc(i32, 6);
    defer allocator.free(buf);

    var grid = Grid(i32).new(buf, 2, 3);

    // Test set and get
    grid.set(0, 0, 1);
    grid.set(0, 1, 2);
    grid.set(0, 2, 3);
    grid.set(1, 0, 4);
    grid.set(1, 1, 5);
    grid.set(1, 2, 6);

    try std.testing.expect(grid.get(0, 0) == 1);
    try std.testing.expect(grid.get(0, 1) == 2);
    try std.testing.expect(grid.get(0, 2) == 3);
    try std.testing.expect(grid.get(1, 0) == 4);
    try std.testing.expect(grid.get(1, 1) == 5);
    try std.testing.expect(grid.get(1, 2) == 6);
}

test "Grid isOob" {
    const allocator = std.testing.allocator;

    const buf = try allocator.alloc(u8, 4);
    defer allocator.free(buf);

    const grid = Grid(u8).new(buf, 2, 2);

    // Test valid positions
    try std.testing.expect(!grid.isOob(0, 0));
    try std.testing.expect(!grid.isOob(0, 1));
    try std.testing.expect(!grid.isOob(1, 0));
    try std.testing.expect(!grid.isOob(1, 1));

    // Test out of bounds positions
    try std.testing.expect(grid.isOob(-1, 0));
    try std.testing.expect(grid.isOob(0, -1));
    try std.testing.expect(grid.isOob(2, 0));
    try std.testing.expect(grid.isOob(0, 2));
    try std.testing.expect(grid.isOob(2, 2));
}

test "Grid newFromInput" {
    const allocator = std.testing.allocator;

    const input = "abc\ndef";
    var grid = try Grid(u8).newFromInput(allocator, input, null);
    defer grid.deinit(allocator);

    try std.testing.expect(grid.rows == 2);
    try std.testing.expect(grid.cols == 3);
    try std.testing.expect(grid.get(0, 0) == 'a');
    try std.testing.expect(grid.get(0, 1) == 'b');
    try std.testing.expect(grid.get(0, 2) == 'c');
    try std.testing.expect(grid.get(1, 0) == 'd');
    try std.testing.expect(grid.get(1, 1) == 'e');
    try std.testing.expect(grid.get(1, 2) == 'f');
}

test "Grid newFromInput with conversion function" {
    const allocator = std.testing.allocator;

    const conv_fn = struct {
        fn convert(c: u8) i32 {
            return @as(i32, c) - '0';
        }
    }.convert;

    const input = "123\n456";
    var grid = try Grid(i32).newFromInput(allocator, input, conv_fn);
    defer grid.deinit(allocator);

    try std.testing.expect(grid.rows == 2);
    try std.testing.expect(grid.cols == 3);
    try std.testing.expect(grid.get(0, 0) == 1);
    try std.testing.expect(grid.get(0, 1) == 2);
    try std.testing.expect(grid.get(0, 2) == 3);
    try std.testing.expect(grid.get(1, 0) == 4);
    try std.testing.expect(grid.get(1, 1) == 5);
    try std.testing.expect(grid.get(1, 2) == 6);
}
