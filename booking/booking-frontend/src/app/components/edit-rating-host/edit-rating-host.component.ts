import { Component, OnInit } from '@angular/core';
import {ActivatedRoute, Params, Router} from "@angular/router";
import {RatingHost} from "../../model/rating-host.model";
import {RatingService} from "../../service/rating.service";
import {User} from "../../model/user.model";
import {UserService} from "../../service/user.service";
import Swal from "sweetalert2";

@Component({
  selector: 'app-edit-rating-host',
  templateUrl: './edit-rating-host.component.html',
  styleUrls: ['./edit-rating-host.component.css']
})
export class EditRatingHostComponent implements OnInit {

  constructor(private ratingService: RatingService,private userService: UserService, private router: Router, private route: ActivatedRoute) { }
  ratingHost: RatingHost = new RatingHost();
  host: User = new User();
  hostName: string = "";
  hostSurname: string = "";
  title = 'Edit rating for host';
  submitted = false;
  stars: number[] = [1, 2, 3, 4, 5];
  selectedValue: number = 0;
  ngOnInit(): void {
    this.route.params.subscribe((params: Params) => {
      this.ratingService.getRatingHostById(params['id']).subscribe(res => {
        console.log(params['id'])
        this.ratingHost = res.ratingHost;
        this.selectedValue = this.ratingHost.rate;
        console.log(this.ratingHost);
        this.userService.getUserById(this.ratingHost.hostId).subscribe(res => {
          this.host = res.user;
          this.hostName = this.host.firstName;
          this.hostSurname = this.host.lastName;
        })
      })
    });


  }
  onSubmit() {
    console.log(this.ratingHost)
    this.ratingHost.rate = this.selectedValue;
    this.ratingService.updateRatingForHost(this.ratingHost).subscribe(
      {
        next: (res) => {
          this.router.navigate(['/ratings-host-by-guest']);
          Swal.fire({
            icon: 'success',
            title: 'Success!',
            text: 'Successfully updated rate for host!',
          })
        },
        error: (e) => {
          console.log(e);
          Swal.fire({
            icon: 'error',
            title: 'Oops...',
            text: 'Something went wrong.',
          })

        }

      });
  }
  countStar(star: number) {
    console.log('Value of star', star);
    this.selectedValue = star;
    console.log(this.selectedValue)
  }
  addClass(star: number) {
    let ab = "";
    for (let i = 0; i < star; i++) {
      ab = "starId" + i;
      // @ts-ignore
      document.getElementById(ab).classList.add("selected");
    }
  }
  removeClass(star: number) {
    let ab = "";
    for (let i = star-1; i >= this.selectedValue; i--) {
      ab = "starId" + i;
      // @ts-ignore
      document.getElementById(ab).classList.remove("selected");
    }
  }

}
