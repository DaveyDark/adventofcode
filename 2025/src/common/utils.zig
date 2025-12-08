const std = @import("std");

pub fn isStringEmpty(s: []const u8) bool {
    if (s.len == 0) {
        return true; // An empty string can be considered as containing only whitespace
    }
    for (s) |char| {
        if (!std.ascii.isWhitespace(char)) {
            return false;
        }
    }
    return true;
}
