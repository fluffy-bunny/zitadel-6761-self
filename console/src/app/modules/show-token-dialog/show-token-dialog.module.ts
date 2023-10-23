import { CommonModule } from '@angular/common';
import { NgModule } from '@angular/core';
import { MatLegacyButtonModule as MatButtonModule } from '@angular/material/legacy-button';
import { MatLegacyTooltipModule as MatTooltipModule } from '@angular/material/legacy-tooltip';
import { TranslateModule } from '@ngx-translate/core';
import { CopyToClipboardModule } from 'src/app/directives/copy-to-clipboard/copy-to-clipboard.module';
import { LocalizedDatePipeModule } from 'src/app/pipes/localized-date-pipe/localized-date-pipe.module';
import { TimestampToDatePipeModule } from 'src/app/pipes/timestamp-to-date-pipe/timestamp-to-date-pipe.module';

import { InfoSectionModule } from '../info-section/info-section.module';
import { ShowTokenDialogComponent } from './show-token-dialog.component';

@NgModule({
  declarations: [ShowTokenDialogComponent],
  imports: [
    CommonModule,
    TranslateModule,
    InfoSectionModule,
    CopyToClipboardModule,
    MatButtonModule,
    MatTooltipModule,
    LocalizedDatePipeModule,
    TimestampToDatePipeModule,
  ],
})
export class ShowTokenDialogModule {}
