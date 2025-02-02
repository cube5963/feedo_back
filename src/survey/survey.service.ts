import { Injectable } from '@nestjs/common';
import { answer } from 'src/db/survey/survey/answer';
import { create } from 'src/db/survey/survey/create';
import { getSurvey } from 'src/db/survey/survey/survey';
import { getAnswer } from 'src/db/survey/survey/getanswer';

@Injectable()
export class SurveyService {
  async Create(survey: any) {
    return await create(survey);
  }

  async Get(surveyid: number) {
    return await getSurvey(surveyid);
  }

  async Answer(answers: any) {
    return await answer(answers);
  }

  async GetAnswer(surveyId: number) {
    return await getAnswer(surveyId);
  }
}
