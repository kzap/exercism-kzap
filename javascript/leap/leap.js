class Year {
  constructor(year) {
    this.year = year;
  }

  isLeap() {
    let year = this.year;

    if (year % 4 === 0 && (year % 100 !== 0 || year % 400 === 0)) {
      return true
    }

    return false;
  }
}

export default Year;
