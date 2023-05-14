import { Injectable } from '@angular/core';
import {HttpClient, HttpHeaders} from "@angular/common/http";
import { Observable } from 'rxjs';
import { Accommodation } from '../model/accommodation.model';
import { Availability } from '../model/availability.model';

@Injectable({
  providedIn: 'root'
})
export class AccommodationService {

  apiHost: string = 'http://localhost:8000/accommodation';
  headers: HttpHeaders = new HttpHeaders({ 'Content-Type': 'application/json', 'Access-Control-Allow-Origin' : '*',
  'Access-Control-Allow-Methods' : 'GET,HEAD,OPTIONS,POST,PUT', 'Access-Control-Allow-Headers' : 'Origin, X-Requested-With, Content-Type, Accept, x-client-key, x-client-token, x-client-secret, Authorization' }  );
  headers2: HttpHeaders = new HttpHeaders({
  'Content-Type': 'application/json',
  });

  constructor(private http: HttpClient) { }

  createAccommodation(accommodation: any) {
    console.log(accommodation);
    return this.http.post<Accommodation>(this.apiHost , accommodation, {headers: this.headers2});
  }

  getAccommodationByHostId(id: any): Observable<any[]> {
    return this.http.get<any[]>(this.apiHost + '/' + id);
  }

  createAvailability(ava: Availability){
    console.log(ava);
    return this.http.post<Availability>(this.apiHost + 'createAvailability', ava);
  }

  getAccommodationById(id: any): Observable<Accommodation> {
    return this.http.get<Accommodation>(this.apiHost + 'getById/' + id);
  }
  searchAccommodations(searchAccommodations: any) {
    console.log(searchAccommodations);
    return this.http.post<Accommodation[]>(this.apiHost + 'search', searchAccommodations);
  }
}
