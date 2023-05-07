export class Flight {
    constructor(
        public id: string = '',
        public departureDateTime: Date = new Date(),
        public arrivalDateTime: Date = new Date(),
        public departure: string = '',
        public arrival: string = '',
        public price: number = 0,
        public totalNumberOfSeats: number = 0,
        public availableSeats: number = 0
      ) {}
}
