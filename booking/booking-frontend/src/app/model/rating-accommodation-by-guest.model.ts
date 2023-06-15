export class RatingAccommodationByGuest {
  constructor(
    public id: string = '',
    public accommodationId: string = '',
    public guestId: string = '',
    public accommodationName: string = '',
    public rate:  number = 0,
    public date: string = '',

  ){}
}
