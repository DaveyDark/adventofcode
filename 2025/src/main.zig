const std = @import("std");
const day1 = @import("day01.zig");

pub fn main() !void {
    var gpa = std.heap.GeneralPurposeAllocator(.{}){};
    try day1.run(gpa.allocator());
}
