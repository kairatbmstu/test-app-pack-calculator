class PackSize {
    constructor(size) {
      this.size = size;
    }
  
    // Getter for size
    getSize() {
      return this.size;
    }
  
    // Method to perform some operation related to pack size
    displayInfo() {
      console.log(`Pack Size: ${this.size}`);
    }
}


class Pack {
    constructor(size, num) {
      this.size = size;
      this.num = num;
    }
  
    // Getter for size
    getSize() {
      return this.size;
    }
  
    // Getter for num
    getNum() {
      return this.num;
    }
  
    // Method to perform some operation related to pack size
    displayInfo() {
      console.log(`Pack Size: ${this.size}, Number: ${this.num}`);
    }
  }