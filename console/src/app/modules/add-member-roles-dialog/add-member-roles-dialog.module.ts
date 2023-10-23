import { CommonModule } from '@angular/common';
import { NgModule } from '@angular/core';
import { MatLegacyButtonModule as MatButtonModule } from '@angular/material/legacy-button';
import { MatLegacyCheckboxModule as MatCheckboxModule } from '@angular/material/legacy-checkbox';
import { MatLegacyTooltipModule as MatTooltipModule } from '@angular/material/legacy-tooltip';
import { TranslateModule } from '@ngx-translate/core';
import { LocalizedDatePipeModule } from 'src/app/pipes/localized-date-pipe/localized-date-pipe.module';
import { RoleTransformPipeModule } from 'src/app/pipes/role-transform/role-transform.module';
import { TimestampToDatePipeModule } from 'src/app/pipes/timestamp-to-date-pipe/timestamp-to-date-pipe.module';

import { AddMemberRolesDialogComponent } from './add-member-roles-dialog.component';

@NgModule({
  declarations: [AddMemberRolesDialogComponent],
  imports: [
    CommonModule,
    TranslateModule,
    MatCheckboxModule,
    MatButtonModule,
    LocalizedDatePipeModule,
    MatTooltipModule,
    RoleTransformPipeModule,
    TimestampToDatePipeModule,
  ],
})
export class AddMemberRolesDialogModule {}
