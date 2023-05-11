export class Availability {
    constructor(
        public startDate: Date = new Date(),
        public endDate: Date = new Date(),
        public accommodationId: string = '',
        public price: number = 0,
        public priceSelection: number = 0,
      ) {}
}
