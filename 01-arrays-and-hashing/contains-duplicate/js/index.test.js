import { containsDuplicate } from "./index"

describe("containsDuplicate", () => {
  it("should return true when there are duplicates", () => {
    expect(containsDuplicate([1,2,3,1])).toBe(true);
  });

  it("should return false when all elements are distinct", () => {
    expect(containsDuplicate([1,2,3,4])).toBe(false);
  });

  it("should return true for multiple duplicates", () => {
    expect(containsDuplicate([1,1,1,3,3,4,3,2,4,2])).toBe(true);
  });
});