package javaassist;

import java.lang.annotation.ElementType;
import java.lang.annotation.Retention;
import java.lang.annotation.RetentionPolicy;
import java.lang.annotation.Target;

@Retention(RetentionPolicy.CLASS)
@Target(ElementType.METHOD)
public @interface MsgProcess {

	Type value() default Type.process;

	public enum Type {
		process(), process_in_procedure();

		public String getName() {
			return process.name();
		}

		public boolean inProcedure() {
			return this == process_in_procedure;
		}
	}
}
