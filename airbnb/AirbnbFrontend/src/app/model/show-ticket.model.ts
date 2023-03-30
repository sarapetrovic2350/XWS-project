export class ShowTicket {
    constructor(
        public id: string = '',
        public dateOfPurchase: string = '',
        public dateOfDeparture: string = '',
        public departure: string = '',
        public arrival: string = '',
        public numberOfTickets: number =0,
        public totalPrice: number =0
      ) {}
}
