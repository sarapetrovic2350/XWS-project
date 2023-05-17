export class Reservation {
    constructor(
        public numberOfGuests: number = 0, 
        public startDate: string = '', 
        public endDate: string = '',
        public userId: string = '', 
        public accommodationId: string = '', 
        public reservationStatus: string = '', 
        public country: string = '', 
        public city: string = '', 
        public name: string = ''
    ){}
}
