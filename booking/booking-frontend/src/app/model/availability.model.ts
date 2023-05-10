export class Availability {
    constructor(
        public startDateTime: Date = new Date(),
        public endlDateTime: Date = new Date(),
        public accommodationId: string = '',
        public price: number = 0,
        public priceSelection: number = 0,
      ) {}
}
