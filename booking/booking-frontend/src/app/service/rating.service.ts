import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { RatingHost } from '../model/rating-host.model';
import { Observable } from 'rxjs';
import { RatingAccommodation } from '../model/rating-accommodation.model';
import {Reservation} from "../model/reservation.model";

@Injectable({
  providedIn: 'root'
})
export class RatingService {

  apiHost: string = 'http://localhost:8000/';
  headers: HttpHeaders = new HttpHeaders({ 'Content-Type': 'application/json', 'Access-Control-Allow-Origin' : '*',
    'Access-Control-Allow-Methods' : 'GET,HEAD,OPTIONS,POST,PUT', 'Access-Control-Allow-Headers' : 'Origin, X-Requested-With, Content-Type, Accept, x-client-key, x-client-token, x-client-secret, Authorization' }  );
  headers2: HttpHeaders = new HttpHeaders({
    'Content-Type': 'application/json',
  });

  constructor(private http: HttpClient) { }

  createRatingForHost(ratingHost: RatingHost): Observable<any> {
    return this.http.post<RatingHost>(this.apiHost + 'createRatingForHost', ratingHost, {headers: this.headers2});
  }

  createRatingForAccommodation(ratingAccommodation: RatingAccommodation): Observable<any> {
    return this.http.post<RatingHost>(this.apiHost + 'createRatingForAccommodation', ratingAccommodation, {headers: this.headers2});
  }
  getRatingsHostByGuestId(guestId: any): Observable<any[]> {
    return this.http.get<any[]>(this.apiHost + 'ratingsHost/' + guestId, {headers: this.headers2});
  }
  getRatingsAccommodationByGuestId(guestId: any): Observable<any[]> {
    return this.http.get<any[]>(this.apiHost + 'ratingsAccommodation/' + guestId, {headers: this.headers2});
  }
  updateRatingForHost(ratingHost: RatingHost): Observable<any> {
    return this.http.post<any>(this.apiHost + 'updateRatingForHost', ratingHost, {headers: this.headers2})
  }
  updateRatingForAccommodation(ratingAccommodation: RatingAccommodation): Observable<any> {
    return this.http.post<any>(this.apiHost + 'updateRatingForAccommodation', ratingAccommodation, {headers: this.headers2})
  }
  getRatingHostById(id: string): Observable<any> {
    return this.http.get<any>(this.apiHost + 'ratingHostById/' + id);
  }
  getRatingAccommodationById(id: string): Observable<any> {
    return this.http.get<any>(this.apiHost + 'ratingAccommodationById/' + id);
  }
  DeleteRatingForHost(id: string) {
    return this.http.post<RatingHost>(this.apiHost + 'deleteRatingForHost/' + id, {headers: this.headers2});
  }
  DeleteRatingForAccommodation(id: string) {
    return this.http.post<RatingHost>(this.apiHost + 'deleteRatingForAccommodation/' + id, {headers: this.headers2});
  }
  getRatingsForHost(hostId : any): Observable<any[]>{
    return this.http.get<any[]>(this.apiHost + 'getRatingsForHost/' + hostId, {headers: this.headers2});
  }
  getAverageRatingForHost(hostId: string): Observable<any> {
    return this.http.get<any>(this.apiHost + 'getAvgRatingForHost/' + hostId, {headers: this.headers2});
  }

}
