const std = @import("std");
const ArrayList = std.ArrayList;

pub fn RangeGroup(comptime T: type) type {
    const Range = [2]T;

    return struct {
        ranges: ArrayList(Range),

        fn compareRange(r1: Range, r2: Range) std.math.Order {
            return std.math.order(r1[0], r2[0]);
        }

        fn rangeOverlap(r1: Range, r2: Range) bool {
            return !(r1[1] < r2[0] or r1[0] > r2[1]);
        }

        fn mergeRanges(r1: Range, r2: Range) Range {
            const start = @min(r1[0], r2[0]);
            const end = @max(r1[1], r2[1]);
            return Range{ start, end };
        }

        pub fn addRange(self: *@This(), allocator: std.mem.Allocator, rng: Range) !void {
            if (self.ranges.items.len == 0) {
                try self.ranges.append(allocator, rng);
                return;
            }
            const pos = std.sort.lowerBound(Range, self.ranges.items, rng, compareRange);
            try self.ranges.insert(allocator, pos, rng);

            while (pos != 0 and pos < self.ranges.items.len and rangeOverlap(self.ranges.items[pos - 1], self.ranges.items[pos])) {
                const r1 = self.ranges.items[pos];
                const r2 = self.ranges.items[pos - 1];
                self.ranges.orderedRemoveMany(&[2]usize{ pos - 1, pos });
                try self.ranges.insert(allocator, pos - 1, mergeRanges(r1, r2));
            }
            while (pos + 1 < self.ranges.items.len and rangeOverlap(self.ranges.items[pos], self.ranges.items[pos + 1])) {
                const r1 = self.ranges.items[pos];
                const r2 = self.ranges.items[pos + 1];
                self.ranges.orderedRemoveMany(&[2]usize{ pos, pos + 1 });
                try self.ranges.insert(allocator, pos, mergeRanges(r1, r2));
            }
        }

        fn compareKey(val: T, range: Range) std.math.Order {
            return std.math.order(val, range[0]);
        }

        pub fn containsValue(self: @This(), val: T) bool {
            var pos = std.sort.upperBound(Range, self.ranges.items, val, compareKey);
            if (pos == 0) return false;
            pos -= 1;
            return val >= self.ranges.items[pos][0] and val <= self.ranges.items[pos][1];
        }

        pub fn newRange() @This() {
            return @This(){ .ranges = ArrayList(Range){} };
        }

        pub fn deinit(self: *@This(), allocator: std.mem.Allocator) void {
            self.ranges.deinit(allocator);
        }
    };
}

test "CombinedRange basic functionality" {
    const allocator = std.testing.allocator;
    const CombRange = RangeGroup(u32);
    const Range = [2]u32;

    var combined = CombRange{
        .ranges = ArrayList([2]u32){},
    };
    defer combined.ranges.deinit(allocator);

    // Test adding a single range
    try combined.addRange(allocator, Range{ 10, 20 });
    try std.testing.expect(combined.ranges.items.len == 1);
    try std.testing.expectEqual(Range{ 10, 20 }, combined.ranges.items[0]);

    // Test adding non-overlapping range
    try combined.addRange(allocator, Range{ 30, 40 });
    try std.testing.expect(combined.ranges.items.len == 2);

    // Test adding overlapping range that should merge
    try combined.addRange(allocator, Range{ 15, 35 });
    try std.testing.expect(combined.ranges.items.len == 1);
    try std.testing.expectEqual(Range{ 10, 40 }, combined.ranges.items[0]);
}

test "CombinedRange range overlap detection" {
    const CombRange = RangeGroup(u32);
    const Range = [2]u32;

    // Test overlapping ranges
    try std.testing.expect(CombRange.rangeOverlap(Range{ 10, 20 }, Range{ 15, 25 }));
    try std.testing.expect(CombRange.rangeOverlap(Range{ 10, 30 }, Range{ 20, 25 }));
    try std.testing.expect(CombRange.rangeOverlap(Range{ 20, 30 }, Range{ 10, 25 }));

    // Test non-overlapping ranges
    try std.testing.expect(!CombRange.rangeOverlap(Range{ 10, 20 }, Range{ 21, 30 }));
    try std.testing.expect(!CombRange.rangeOverlap(Range{ 30, 40 }, Range{ 10, 20 }));

    // Test adjacent ranges (should not overlap)
    try std.testing.expect(!CombRange.rangeOverlap(Range{ 10, 20 }, Range{ 21, 30 }));
}

test "CombinedRange merge ranges" {
    const CombRange = RangeGroup(u32);
    const Range = [2]u32;

    // Test merging overlapping ranges
    try std.testing.expectEqual(Range{ 10, 30 }, CombRange.mergeRanges(Range{ 10, 20 }, Range{ 15, 30 }));
    try std.testing.expectEqual(Range{ 5, 25 }, CombRange.mergeRanges(Range{ 10, 20 }, Range{ 5, 25 }));
    try std.testing.expectEqual(Range{ 10, 30 }, CombRange.mergeRanges(Range{ 15, 30 }, Range{ 10, 20 }));
}

test "CombinedRange contains value" {
    const allocator = std.testing.allocator;
    const CombRange = RangeGroup(u32);
    const Range = [2]u32;

    var combined = CombRange{
        .ranges = ArrayList([2]u32){},
    };
    defer combined.ranges.deinit(allocator);

    try combined.addRange(allocator, Range{ 10, 20 });
    try combined.addRange(allocator, Range{ 30, 40 });

    // Test values within ranges
    try std.testing.expect(combined.containsValue(15));
    try std.testing.expect(combined.containsValue(35));
    try std.testing.expect(combined.containsValue(10)); // boundary
    try std.testing.expect(combined.containsValue(20)); // boundary

    // Test values outside ranges
    try std.testing.expect(!combined.containsValue(5));
    try std.testing.expect(!combined.containsValue(25));
    try std.testing.expect(!combined.containsValue(45));
}
