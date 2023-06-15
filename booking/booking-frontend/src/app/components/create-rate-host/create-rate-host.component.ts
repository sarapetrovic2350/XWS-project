import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Params, Router } from '@angular/router';
import { RatingHost } from 'src/app/model/rating-host.model';
import { User } from 'src/app/model/user.model';
import { RatingService } from 'src/app/service/rating.service';
import { UserService } from 'src/app/service/user.service';
import Swal from 'sweetalert2';

@Component({
  selector: 'app-create-rate-host',
  templateUrl: './create-rate-host.component.html',
  styleUrls: ['./create-rate-host.component.css']
})
export class CreateRateHostComponent implements OnInit {

  constructor( private router: Router,
    private userService : UserService,
    private route: ActivatedRoute,
    private ratinService: RatingService) { }

  title = 'Rate host';
  submitted = false;

  user: User = new User();
  guest: User = new User();
  ratingHost: RatingHost = new RatingHost();


  ngOnInit(): void {

    this.route.params.subscribe((params: Params) => {
      this.userService.getUserByEmail(params['email']).subscribe(res => {
        console.log(params['email'])
        this.user = res.user;
        console.log(res);
      })
    });

    let userRole = this.userService.getLoggedInUserRole()
    let userEmail = this.userService.getLoggedInUserEmail()
    this.userService.getUserByEmail(userEmail).subscribe(res => {
      this.guest = res.user;
      console.log(res.user.id)
    })

  }

  onSubmit(){

    this.ratingHost.guestId = this.guest.id;
    this.ratingHost.hostId = this.user.id;
    console.log(this.ratingHost)

    this.ratinService.createRatingForHost(this.ratingHost).subscribe(
      {
        next: (res) => {
          this.router.navigate(['/ratings-host-by-guest']);
          Swal.fire({
            icon: 'success',
            title: 'Success!',
            text: 'Successfully rated host!',
          })

        },
        error: (e) => {
          this.submitted = false;
          console.log(e);
          Swal.fire({
            icon: 'error',
            title: 'Oops...',
            text: 'Guest does not have a reservation in past that is not canceled.',
          })

        }

      });




  }

}
