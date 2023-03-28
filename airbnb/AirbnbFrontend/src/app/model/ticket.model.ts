export class Ticket {
    constructor(
        public id: string = '',
        public numberOfTickets: number =0, 
        public idUser: string ='',
        public idFlight: string = ''
      ) {}
}
