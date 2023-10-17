package javaassist;

import java.util.HashMap;
import java.util.Map;

import javaassist.MsgProcess.Type;
import javassist.ByteArrayClassPath;
import javassist.CannotCompileException;
import javassist.ClassPool;
import javassist.CtClass;
import javassist.CtMethod;
import javassist.CtNewMethod;

public class Replace {

    private static ClassPool cp;

    private static Map<String, Class<?>> classMap = new HashMap<String, Class<?>>();
    
    public static void main(String[] args) throws Exception {
        cp = ClassPool.getDefault();
        cp.insertClassPath(new javassist.LoaderClassPath(Replace.class.getClassLoader()));

        // 修改CTest的process方法体
        String className = "javaassist.MsgTest";
        CtClass c = getClass(className);
        for (CtMethod method : c.getDeclaredMethods()) {
			Object annotation = method.getAnnotation(MsgProcess.class);
			if (annotation == null) {
				continue;
			}

			CtClass[] parameterTypes = method.getParameterTypes();
			if (parameterTypes.length != 1) {
				System.err.println("method startwith exex , but is not gen source :" + method.getName());
				continue;
			}

			Type value = MsgProcess.class.cast(annotation).value();
			String methodName = value.getName();
			CtClass msgClass = genClass(className, method.getName(), parameterTypes[0].getName(), methodName, value.inProcedure());
 
			Class<?> class1;
			try {
				class1 = msgClass.toClass();
			} catch (Exception e) {
				throw new RuntimeException("to class fail :" + className + " msg :" + msgClass.getName());
			}
            classMap.put(msgClass.getName(), class1);
		}
		c.detach();
        
        new CTest().process();
    }

    private static CtClass genClass(String processClassName, String proessMethodName, String msgClassName, String msgMethodName, boolean inProcedure) throws Exception {
		CtClass msgClass = getClass(msgClassName);
		renameMethod(msgMethodName, msgClass);

		msgClass.getClassPool().clearImportedPackages();
		msgClass.getClassPool().importPackage(processClassName);
		CtMethod make = CtNewMethod.make("public void " + msgMethodName + "()throws Exception { }", msgClass);
		msgClass.addMethod(make);
		if (inProcedure) {
			String name = getProcedureName(msgClassName);
			createProcedure(name, msgClassName, processClassName, proessMethodName);
			make.setBody("new " + name + "($0).queueMsg($0);", processClassName, proessMethodName);
			return msgClass;
		}
		try {
			make.setBody("try { $proceed($0);} catch (Exception e) {" + msgMethodName + "old" + "(); }", processClassName, proessMethodName);
		} catch (Exception e) { // javassist bug 包名与类名相同时,编译失败.
			String[] split = processClassName.split("\\.");
			make.setBody("try { $proceed($0);} catch (Exception e) {" + msgMethodName + "old" + "(); }", split[split.length - 1], proessMethodName);
		}
		return msgClass;
	}

    private static String getProcedureName(String msgClassName) {
		String[] split = msgClassName.split("\\.");
		StringBuilder sb = new StringBuilder();
		for (int i = 0; i < split.length; i++) {
			sb.append(split[i]);
			if (i < split.length - 1) {
				sb.append(".");
				if (i == split.length - 2) {
					sb.append("P");
				}
			}
		}
		return sb.toString();
	}


    private static void renameMethod(String methodName, CtClass msgClass) throws CannotCompileException {
		try {
			CtMethod oldMethod = msgClass.getDeclaredMethod(methodName);
			oldMethod.setName(methodName + "old");
		} catch (Exception e) {
			CtMethod make = CtNewMethod.make("public void " + methodName + "old(){ }", msgClass);
			msgClass.addMethod(make);
		}
	}

    static void createProcedure(String name, String msgClassName, String processClassName, String proessMethodName) throws Exception {
		CtClass proc = getClass("javaassist.TempProcedure");
		proc.setName(name);
		CtMethod m = proc.getDeclaredMethod("process");

		CtClass processClass = getClass(processClassName);
		CtMethod proessMethod = processClass.getDeclaredMethod(proessMethodName, new CtClass[]{getClass(msgClassName)});
		CtClass returnType = proessMethod.getReturnType();
		try {
			if (returnType.getName().equals("boolean")) {
				m.setBody("return $proceed((" + msgClassName + ")msg);", processClassName, proessMethodName);
			} else {
				m.setBody("{$proceed((" + msgClassName + ")msg);return true;}", processClassName, proessMethodName);
			}
		} catch (Exception e) { // javassist bug 包名与类名相同时,编译失败.
			String[] split = processClassName.split("\\.");
			if (returnType.getName().equals("boolean")) {
				m.setBody("return $proceed((" + msgClassName + ")msg);", split[split.length - 1], proessMethodName);
			} else {
				m.setBody("{$proceed((" + msgClassName + ")msg);return true;}", split[split.length - 1], proessMethodName);
			}
		}
		cp.insertClassPath(new ByteArrayClassPath(name, proc.toBytecode()));
        classMap.put(name, proc.toClass());
	}

    private static CtClass getClass(String className) {
		CtClass clazz = null;
		try {
			clazz = cp.get(className);
		} catch (Exception e) {
		}
		return clazz;
	}
}
