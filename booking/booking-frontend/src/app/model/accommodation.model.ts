import {Address} from "./address.model";

export class Accommodation {
    constructor(
        public id: string = '',
        public name: string = '',
        public minNumberOfGuests: number = 0,
        public maxNumberOfGuests: number = 0,
        public hostId: string = '',
        public benefits: string[] = [],
        public address: Address = new Address()
      ) {}
}
