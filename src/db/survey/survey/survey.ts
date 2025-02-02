import prisma from 'src/db/db.service';

export async function surveyCreate(
  title: string,
  description: string,
  loginid: string,
): Promise<number> {
  const userid = await prisma.user.findUnique({
    where: {
      loginid: loginid,
    },
    select: {
      id: true,
    },
  });

  const survey = await prisma.survey.create({
    data: {
      title: title,
      description: description,
      userId: userid.id,
    },
  });
  return survey.id;
}

export async function getSurvey(surveyid: number) {
  const survey = await prisma.survey.findUnique({
    where: {
      id: surveyid,
    },
    include: {
      questions: true,
    },
  });

  const response = {
    id: survey.id,
    title: survey.title,
    description: survey.description,
    questions: survey.questions.map((q) => ({
      id: q.id,
      title: q.title,
      type: q.type,
      content:
        typeof q.content === 'string' ? JSON.parse(q.content) : q.content,
    })),
  };

  return response;
}
