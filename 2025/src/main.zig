const std = @import("std");
const day1 = @import("day01.zig");
const day2 = @import("day02.zig");
const day3 = @import("day03.zig");

pub fn main() !void {
    var gpa = std.heap.GeneralPurposeAllocator(.{}){};
    try day1.run(gpa.allocator());
    try day2.run(gpa.allocator());
    try day3.run(gpa.allocator());
}
