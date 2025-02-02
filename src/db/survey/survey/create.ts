import { QuestionModule, QuestionType } from '../module/question';
import { surveyCreate } from './survey';

export async function create(survey: {
  title: string;
  description: string;
  userId: string;
  questions: { title: string; type: string; content: any }[];
}) {
  const surveyid = await surveyCreate(
    survey.title,
    survey.description,
    survey.userId,
  );

  const questions = survey.questions.map(
    (q, index) =>
      new QuestionModule(
        surveyid,
        q.title,
        index + 1,
        q.type as unknown as QuestionType,
        q.content,
      ),
  );

  console.log(questions);
}
