import { ComponentFixture, TestBed } from '@angular/core/testing';

import { AddItemViewComponent } from './add-item-view.component';

describe('AddItemViewComponent', () => {
  let component: AddItemViewComponent;
  let fixture: ComponentFixture<AddItemViewComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [AddItemViewComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(AddItemViewComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
