class Year {
  constructor(year) {
    this.year = year;
  }

  isLeap() {
    let isLeapYear = false;
    let year = this.year;

    if (year % 4 === 0) {
      isLeapYear = true;

      if (year % 100 === 0) {
        isLeapYear = false;

        if (year % 400 === 0) {
          isLeapYear = true;
        }
      }
    }

    return isLeapYear;
  }
}

export default Year;
