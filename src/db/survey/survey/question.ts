import { QuestionType } from '../module/question';
import prisma from 'src/db/db.service';

export async function questionCreate(
  title: string,
  surveyid: number,
  order: number,
  type: QuestionType,
  content: any,
): Promise<void> {
  await prisma.question.create({
    data: {
      title: title,
      surveyId: surveyid,
      order: order,
      type: type,
      content: content,
    },
  });
}
