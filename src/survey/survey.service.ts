import { Injectable } from '@nestjs/common';
import { create } from 'src/db/survey/survey/create';

@Injectable()
export class SurveyService {
  async Create(survey: any) {
    await create(survey);
  }
}
