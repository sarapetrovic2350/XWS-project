export class RatingAccommodationForHost {
  constructor(
    public id: string = '',
    public accommodationId: string = '',
    public accommodationName: string = '',
    public guestId: string = '',
    public guestName: string = '',
    public guestSurname: string = '',
    public rate:  number = 0,
    public date: string = '',
  ){}
}
