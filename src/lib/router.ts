import { initTRPC } from "@trpc/server";
import { wrap } from "@decs/typeschema";
import { string } from "valibot";

const t = initTRPC.create();

export const appRouter = t.router({
  hello: t.procedure.input(wrap(string())).query(({ input }) => {
    return `hello ${input ?? "world"}`;
  })
});

export type AppRouter = typeof appRouter;