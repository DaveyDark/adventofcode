const std = @import("std");

pub fn Queue(comptime T: type) type {
    return struct {
        buf: []T,
        head: usize = 0,
        len: usize = 0,
        allocator: std.mem.Allocator,

        pub fn init(allocator: std.mem.Allocator, cap: usize) !@This() {
            return .{
                .buf = try allocator.alloc(T, cap),
                .allocator = allocator,
            };
        }

        pub fn deinit(self: *@This()) void {
            self.allocator.free(self.buf);
        }

        fn grow(self: *@This()) !void {
            const new_cap = self.buf.len * 2 + 1;
            var new_buf = try self.allocator.alloc(T, new_cap);

            // copy in logical order
            for (0..self.len) |i| {
                new_buf[i] = self.buf[(self.head + i) % self.buf.len];
            }

            self.allocator.free(self.buf);
            self.buf = new_buf;
            self.head = 0;
        }

        pub fn push(self: *@This(), v: T) !void {
            if (self.len == self.buf.len) try self.grow();
            const tail = (self.head + self.len) % self.buf.len;
            self.buf[tail] = v;
            self.len += 1;
        }

        pub fn pop(self: *@This()) ?T {
            if (self.len == 0) return null;
            const v = self.buf[self.head];
            self.head = (self.head + 1) % self.buf.len;
            self.len -= 1;
            return v;
        }

        pub fn empty(self: *@This()) bool {
            return self.len == 0;
        }
    };
}

test "queue FIFO order" {
    var q = try Queue(i32).init(std.testing.allocator, 2);
    defer q.deinit();

    try q.push(1);
    try q.push(2);
    try q.push(3); // forces grow

    try std.testing.expectEqual(@as(?i32, 1), q.pop());
    try std.testing.expectEqual(@as(?i32, 2), q.pop());
    try std.testing.expectEqual(@as(?i32, 3), q.pop());
    try std.testing.expectEqual(@as(?i32, null), q.pop());
}

test "queue wraparound" {
    var q = try Queue(i32).init(std.testing.allocator, 2);
    defer q.deinit();

    try q.push(10);
    try q.push(20);
    _ = q.pop(); // remove 10
    try q.push(30); // wraps internally

    try std.testing.expectEqual(@as(?i32, 20), q.pop());
    try std.testing.expectEqual(@as(?i32, 30), q.pop());
}
