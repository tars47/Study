class Test {
  test() {
    console.log("in test method");
    console.log(Test.map);
    return;
  }

  static map = {
    a: "1",
    b: "2",
  };
}

module.exports = new Test();
