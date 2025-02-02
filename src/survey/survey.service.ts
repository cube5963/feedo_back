import { Injectable } from '@nestjs/common';
import { create } from 'src/db/survey/survey/create';
import { getSurvey } from 'src/db/survey/survey/survey';

@Injectable()
export class SurveyService {
  async Create(survey: any) {
    return await create(survey);
  }

  async Get(surveyid: number) {
    return await getSurvey(surveyid);
  }
}
