export class Reservation {
    constructor(
        numberOfGuests: number = 0, 
        startDate: Date = new Date(), 
        endDate: Date = new Date(),
        userId: string = '', 
        accommodationId: string = '', 
        reservationStatus: string = ''
    ){}
}
