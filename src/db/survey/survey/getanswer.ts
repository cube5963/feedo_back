import prisma from 'src/db/db.service';

export async function getAnswer(surveyId: number): Promise<any> {
  return await prisma.survey.findUnique({
    where: { id: surveyId },
    include: {
      questions: {
        include: {
          responses: true,
        },
      },
    },
  });
}
