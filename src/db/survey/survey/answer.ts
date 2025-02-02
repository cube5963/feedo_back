import prisma from 'src/db/db.service';
import { AnswersInput } from 'src/survey/dto/answer.input';

export async function answer(answers: AnswersInput): Promise<boolean> {
  await prisma.response.createMany({
    data: answers.answers.map((q) => ({
      surveyId: answers.surveyId,
      questionId: q.questionId,
      data: JSON.parse(q.answer),
    })),
  });

  return true;
}
