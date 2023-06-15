export class RatingHostByGuest {
  constructor(
    public id: string = '',
    public hostId: string = '',
    public guestId: string = '',
    public hostName: string = '',
    public hostSurname: string = '',
    public rate:  number = 0,
    public date: string = '',
    
  ){}
}
