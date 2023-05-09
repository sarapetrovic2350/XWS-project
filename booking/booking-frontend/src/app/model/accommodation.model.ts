import {Address} from "./address.model";

export class Accommodation {
    constructor(
        public Name: string = '',
        public MinNumberOfGuests: number = 0,
        public MaxNumberOfGuests: number = 0,
        public HostId: string = '',
        public Benefits: string[] = [],
        public address: Address = new Address()
      ) {}
}
