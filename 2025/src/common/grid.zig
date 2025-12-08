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
