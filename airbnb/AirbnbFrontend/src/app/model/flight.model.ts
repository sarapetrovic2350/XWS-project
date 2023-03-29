export class Flight {
    constructor(
        public id: string = '',
        public departureDate: string = '',
        public departureTime: string = '',
        public arrivalDate: string = '',
        public arrivalTime: string = '',
        public departure: string = '',
        public arrival: string = '',
        public price: number = 0,
        public totalNumberOfSeats: number = 0, 
        public availableSeats: number = 0
      ) {}
}
