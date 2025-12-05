const std = @import("std");

test "import all src" {
    _ = @import("day01.zig");
    _ = @import("day02.zig");
    _ = @import("day03.zig");
}
