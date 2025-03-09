import { bootstrapApplication } from '@angular/platform-browser';
import { appConfig } from './app/app.config';
import { AppComponent } from './app/app.component';
import {provideRouter} from '@angular/router';
import {HomeViewComponent} from './app/views/home-view/home-view.component';
import {InventoryViewComponent} from './app/views/inventory-view/inventory-view.component';

bootstrapApplication(AppComponent, {
  providers: [
    provideRouter([
      { path: '', component: HomeViewComponent },
      { path: 'inventory', component: InventoryViewComponent }
    ])
  ]
}).catch(err => console.error(err));
