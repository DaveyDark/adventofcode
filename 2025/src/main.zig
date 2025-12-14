const std = @import("std");
const day1 = @import("day01.zig");
const day2 = @import("day02.zig");
const day3 = @import("day03.zig");
const day4 = @import("day04.zig");
const day5 = @import("day05.zig");
const day6 = @import("day06.zig");
const day7 = @import("day07.zig");
const day8 = @import("day08.zig");
const day9 = @import("day09.zig");
const day10 = @import("day10.zig");

pub fn main() !void {
    var gpa = std.heap.GeneralPurposeAllocator(.{}){};
    try day1.run(gpa.allocator());
    try day2.run(gpa.allocator());
    try day3.run(gpa.allocator());
    try day4.run(gpa.allocator());
    try day5.run(gpa.allocator());
    try day6.run(gpa.allocator());
    try day7.run(gpa.allocator());
    try day8.run(gpa.allocator());
    try day10.run(gpa.allocator());
}
