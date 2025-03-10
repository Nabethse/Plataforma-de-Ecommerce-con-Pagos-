import { Component } from '@angular/core';
import { OrderService } from '../app/app.component';

@Component({
  selector: 'app-get-order',
  templateUrl: './get-order.component.html',
  styleUrls: ['./get-order.component.css']
})
export class GetOrderComponent {
  orderId: string = '';
  order: any;

  constructor(private orderService: OrderService) {}

  getOrder() {
    this.orderService.getOrder(this.orderId).subscribe(
      (res) => {
        this.order = res;
      },
      (err) => {
        console.error(err);
      }
    );
  }
}