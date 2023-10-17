package javaassist;

import javaassist.MsgProcess.Type;

public class MsgTest {

    @MsgProcess(Type.process_in_procedure)
    public static void process(final CTest msg) {
        System.out.println("process CTest");
    }
}
